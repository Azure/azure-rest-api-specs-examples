from azure.identity import DefaultAzureCredential
from azure.mgmt.testbase import TestBase

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-testbase
# USAGE
    python package_create.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = TestBase(
        credential=DefaultAzureCredential(),
        subscription_id="subscription-id",
    )

    response = client.packages.begin_create(
        resource_group_name="contoso-rg1",
        test_base_account_name="contoso-testBaseAccount1",
        package_name="contoso-package2",
        parameters={
            "location": "westus",
            "properties": {
                "applicationName": "contoso-package2",
                "blobPath": "storageAccountPath/package.zip",
                "flightingRing": "Insider Beta Channel",
                "targetOSList": [
                    {"osUpdateType": "Security updates", "targetOSs": ["Windows 10 2004", "Windows 10 1903"]}
                ],
                "tests": [
                    {
                        "commands": [
                            {
                                "action": "Install",
                                "alwaysRun": True,
                                "applyUpdateBefore": False,
                                "content": "app/scripts/install/job.ps1",
                                "contentType": "Path",
                                "maxRunTime": 1800,
                                "name": "Install",
                                "restartAfter": True,
                                "runAsInteractive": True,
                                "runElevated": True,
                            },
                            {
                                "action": "Launch",
                                "alwaysRun": False,
                                "applyUpdateBefore": True,
                                "content": "app/scripts/launch/job.ps1",
                                "contentType": "Path",
                                "maxRunTime": 1800,
                                "name": "Launch",
                                "restartAfter": False,
                                "runAsInteractive": True,
                                "runElevated": True,
                            },
                            {
                                "action": "Close",
                                "alwaysRun": False,
                                "applyUpdateBefore": False,
                                "content": "app/scripts/close/job.ps1",
                                "contentType": "Path",
                                "maxRunTime": 1800,
                                "name": "Close",
                                "restartAfter": False,
                                "runAsInteractive": True,
                                "runElevated": True,
                            },
                            {
                                "action": "Uninstall",
                                "alwaysRun": True,
                                "applyUpdateBefore": False,
                                "content": "app/scripts/uninstall/job.ps1",
                                "contentType": "Path",
                                "maxRunTime": 1800,
                                "name": "Uninstall",
                                "restartAfter": False,
                                "runAsInteractive": True,
                                "runElevated": True,
                            },
                        ],
                        "isActive": True,
                        "testType": "OutOfBoxTest",
                    }
                ],
                "version": "1.0.0",
            },
            "tags": {},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/testbase/resource-manager/Microsoft.TestBase/preview/2022-04-01-preview/examples/PackageCreate.json
if __name__ == "__main__":
    main()
