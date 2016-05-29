package template

var internal_build = `# {{.Update}}
set -eu
# oas build context
OAS_CURRENT_TAG=$(git tag | tail -n1)
OAS_CURRENT_TAG=${OAS_CURRENT_TAG:-0.SNAPSHOT}
OAS_VERSION_MAJOR_MINOR="${DRONE_TAG:-$OAS_CURRENT_TAG}"
OAS_VERSION_RELEASE="${DRONE_BUILD_NUMBER:-$(date +%s)}"
OAS_VERSION="${OAS_VERSION_MAJOR_MINOR}.${OAS_VERSION_RELEASE}"
OAS_PACKAGE_NAME="{{.Project.Project_name}}"

export OAS_VERSION
export OAS_PACKAGE_NAME

rm -rv target-root
mkdir -pv root
cp -rv root target-root

source scripts/build

# build the package
rm -rv target
mkdir -p target
if ! which fpm > /dev/null 2>&1
then
  echo fpm no encontrado, tratando de instalar
  gem install fpm
fi

if fpm --version
then
  fpm --package=target -C target-root -s dir -t rpm --name="${OAS_PACKAGE_NAME}" --version="${OAS_VERSION}" --before-install scripts/before-install --after-install scripts/after-install --before-remove scripts/before-remove --after-remove scripts/after-remove --before-upgrade scripts/before-upgrade --after-upgrade scripts/after-upgrade --rpm-os linux .
else
  echo no se encontró fpm
  exit 1
fi
`
