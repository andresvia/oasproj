package template

var internal_build = `
set -eu
# oas build context
OAS_CURRENT_TAG=$(git tag | tail -n1)
OAS_CURRENT_TAG=${OAS_CURRENT_TAG:-0.SNAPSHOT}
OAS_VERSION_MAJOR_MINOR="${DRONE_TAG:-$OAS_CURRENT_TAG}"
OAS_VERSION_RELEASE="${DRONE_BUILD_NUMBER:-$(date +%s)}"
OAS_VERSION="${OAS_VERSION_MAJOR_MINOR}.${OAS_VERSION_RELEASE}"
OAS_PACKAGE_NAME="{{.Project_name}}"

export OAS_VERSION
export OAS_PACKAGE_NAME

source scripts/build

# build the package
rm -rf target
mkdir -p target
fpm --package=target -C root -s dir -t rpm --name="${OAS_PACKAGE_NAME}" --version="${OAS_VERSION}" --before-install scripts/before-install --after-install scripts/after-install --before-remove scripts/before-remove --after-remove scripts/after-remove --rpm-os linux .
`
