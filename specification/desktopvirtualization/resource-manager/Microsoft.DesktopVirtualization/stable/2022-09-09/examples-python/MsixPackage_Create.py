from azure.identity import DefaultAzureCredential
from azure.mgmt.desktopvirtualization import DesktopVirtualizationMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-desktopvirtualization
# USAGE
    python msix_package_create.py

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

    response = client.msix_packages.create_or_update(
        resource_group_name="resourceGroup1",
        host_pool_name="hostpool1",
        msix_package_full_name="msixpackagefullname",
        msix_package={
            "properties": {
                "displayName": "displayname",
                "imagePath": "imagepath",
                "isActive": False,
                "isRegularRegistration": False,
                "lastUpdated": "2008-09-22T14:01:54.9571247Z",
                "packageApplications": [
                    {
                        "appId": "ApplicationId",
                        "appUserModelID": "AppUserModelId",
                        "description": "application-desc",
                        "friendlyName": "friendlyname",
                        "iconImageName": "Apptile",
                        "rawIcon": "VGhpcyBpcyBhIHN0cmluZyB0byBoYXNo",
                        "rawPng": "VGhpcyBpcyBhIHN0cmluZyB0byBoYXNo",
                    }
                ],
                "packageDependencies": [
                    {
                        "dependencyName": "MsixTest_Dependency_Name",
                        "minVersion": "version",
                        "publisher": "PublishedName",
                    }
                ],
                "packageFamilyName": "MsixPackage_FamilyName",
                "packageName": "MsixPackage_name",
                "packageRelativePath": "packagerelativepath",
                "version": "version",
            }
        },
    )
    print(response)


# x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2022-09-09/examples/MsixPackage_Create.json
if __name__ == "__main__":
    main()
