package template

var internal_build = `# {{.Update}}
set -eu
# detect os
if [ "$(uname -sm)" = "Linux x86_64" ]
then
  GO_BUILD="go build -v"
  GO_ARTIFACT="{{.Project.Project_name}}"
else
  GO_BUILD="gox -os Linux -arch amd64"
  GO_ARTIFACT="{{.Project.Project_name}}_linux_amd64"
fi
# oas build context
OAS_CURRENT_TAG="$(git tag | tail -n1)"
IFS="." read -r OAS_VERSION_MAJOR OAS_VERSION_MINOR OAS_VERSION_PATCH _JUNK <<< "${OAS_CURRENT_TAG}"
OAS_VERSION_MAJOR="${OAS_VERSION_MAJOR/v/}"
OAS_VERSION_MAJOR="${OAS_VERSION_MAJOR:-0}"
OAS_VERSION_MINOR="${OAS_VERSION_MINOR:-0}"
OAS_VERSION_PATCH="${OAS_VERSION_PATCH:-0}"
if [ "${DRONE:-false}" = "true" ]
then
  OAS_VERSION_META="drone.${DRONE_BUILD_NUMBER:-missing_drone_build_number.$(date +%s)}"
else
  OAS_VERSION_META="local.$(date +%s)"
fi
OAS_VERSION="${OAS_VERSION_MAJOR}.${OAS_VERSION_MINOR}.${OAS_VERSION_PATCH}-${OAS_VERSION_META}"
OAS_NAME="{{.Project.Project_name}}"

export OAS_VERSION
export OAS_NAME
export GO_BUILD
export GO_ARTIFACT

rm -rvf target-root
mkdir -pv root
cp -rv root target-root

source scripts/build

# build the package
rm -rvf target
mkdir -p target
if ! which fpm > /dev/null 2>&1
then
  echo fpm no encontrado, tratando de instalar
  gem install fpm
fi

if fpm --help > /dev/null
then
  set -x
  fpm --description '{{.Project.Project_description}}' --package target -C target-root -s dir -t rpm --name "{{.Project.Project_name}}" {{range .Project.Package_dependencies}} -d {{.}} {{end}} --version "${OAS_VERSION}" --before-install scripts/before-install --after-install scripts/after-install --before-remove scripts/before-remove --after-remove scripts/after-remove --before-upgrade scripts/before-upgrade --after-upgrade scripts/after-upgrade --rpm-os linux {{if .Project.Sign_package}}--rpm-sign{{end}} .
else
  echo no se encontró fpm
  exit 1
fi
`
