from azure.identity import DefaultAzureCredential

from azure.mgmt.desktopvirtualization import DesktopVirtualizationMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-desktopvirtualization
# USAGE
    python app_attach_package_create.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DesktopVirtualizationMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="daefabc0-95b4-48b3-b645-8a753a63c4fa",
    )

    response = client.app_attach_package.create_or_update(
        resource_group_name="resourceGroup1",
        app_attach_package_name="msixpackagefullname",
        app_attach_package={
            "location": "southcentralus",
            "properties": {
                "failHealthCheckOnStagingFailure": "NeedsAssistance",
                "hostPoolReferences": [],
                "image": {
                    "certificateExpiry": "2023-01-02T17:18:19.1234567Z",
                    "certificateName": "certName",
                    "displayName": "displayname",
                    "imagePath": "imagepath",
                    "isActive": False,
                    "isRegularRegistration": False,
                    "lastUpdated": "2008-09-22T14:01:54.9571247Z",
                    "packageAlias": "msixpackagealias",
                    "packageApplications": [
                        {
                            "appId": "AppId",
                            "appUserModelID": "AppUserModelId",
                            "description": "PackageApplicationDescription",
                            "friendlyName": "FriendlyName",
                            "iconImageName": "Iconimagename",
                            "rawIcon": "VGhpcyBpcyBhIHN0cmluZyB0byBoYXNo",
                            "rawPng": "VGhpcyBpcyBhIHN0cmluZyB0byBoYXNo",
                        }
                    ],
                    "packageDependencies": [
                        {
                            "dependencyName": "MsixPackage_Dependency_Name",
                            "minVersion": "packageDep_version",
                            "publisher": "MsixPackage_Dependency_Publisher",
                        }
                    ],
                    "packageFamilyName": "MsixPackage_FamilyName",
                    "packageFullName": "MsixPackage_FullName",
                    "packageName": "MsixPackageName",
                    "packageRelativePath": "packagerelativepath",
                    "version": "packageversion",
                },
                "keyVaultURL": "",
            },
        },
    )
    print(response)


# x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2024-04-03/examples/AppAttachPackage_Create.json
if __name__ == "__main__":
    main()
