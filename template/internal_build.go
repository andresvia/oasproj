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
OAS_CURRENT_TAG=$(git tag | tail -n1)
OAS_CURRENT_TAG=${OAS_CURRENT_TAG:-0.SNAPSHOT}
OAS_VERSION_MAJOR_MINOR="${DRONE_TAG:-$OAS_CURRENT_TAG}"
OAS_VERSION_RELEASE="${DRONE_BUILD_NUMBER:-$(date +%s)}"
OAS_VERSION="${OAS_VERSION_MAJOR_MINOR}.${OAS_VERSION_RELEASE}"
OAS_NAME="{{.Project.Project_name}}"
OAS_SIGN_PACKAGE="${OAS_SIGN_PACKAGE:-${DRONE_TAG:-}}"

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

if fpm --version
then
  fpm_extra_flags=""
  if [ -n "${OAS_SIGN_PACKAGE}" ]
  then
    fpm_extra_flags="${fpm_extra_flags} --rpm-sign"
  fi
  fpm --package=target -C target-root -s dir -t rpm --name="${OAS_NAME}" --version="${OAS_VERSION}" --before-install scripts/before-install --after-install scripts/after-install --before-remove scripts/before-remove --after-remove scripts/after-remove --before-upgrade scripts/before-upgrade --after-upgrade scripts/after-upgrade --rpm-os linux ${fpm_extra_flags} .
else
  echo no se encontró fpm
  exit 1
fi
`
