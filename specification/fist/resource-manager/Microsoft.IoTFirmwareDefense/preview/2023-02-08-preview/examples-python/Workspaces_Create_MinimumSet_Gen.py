from azure.identity import DefaultAzureCredential
from azure.mgmt.iotfirmwaredefense import IoTFirmwareDefenseMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-iotfirmwaredefense
# USAGE
    python workspaces_create_minimum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = IoTFirmwareDefenseMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="5443A01A-5242-4950-AC1A-2DD362180254",
    )

    response = client.workspaces.create(
        resource_group_name="rgworkspaces",
        workspace_name="E___-3",
        workspace={"location": "jjwbseilitjgdrhbvvkwviqj"},
    )
    print(response)


# x-ms-original-file: specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/preview/2023-02-08-preview/examples/Workspaces_Create_MinimumSet_Gen.json
if __name__ == "__main__":
    main()
