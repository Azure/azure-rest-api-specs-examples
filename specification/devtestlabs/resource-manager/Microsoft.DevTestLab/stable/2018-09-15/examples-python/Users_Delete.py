from azure.identity import DefaultAzureCredential

from azure.mgmt.devtestlabs import DevTestLabsClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-devtestlabs
# USAGE
    python users_delete.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DevTestLabsClient(
        credential=DefaultAzureCredential(),
        subscription_id="{subscriptionId}",
    )

    client.users.begin_delete(
        resource_group_name="resourceGroupName",
        lab_name="{devtestlabName}",
        name="{userName}",
    ).result()


# x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Users_Delete.json
if __name__ == "__main__":
    main()
