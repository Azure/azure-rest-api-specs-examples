from azure.identity import DefaultAzureCredential
from azure.mgmt.devhub import DevHubMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-devhub
# USAGE
    python workflow_update_tags.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DevHubMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="subscriptionId1",
    )

    response = client.workflow.update_tags(
        resource_group_name="resourceGroup1",
        workflow_name="workflow1",
        parameters={"tags": {"promote": "false", "resourceEnv": "testing"}},
    )
    print(response)


# x-ms-original-file: specification/developerhub/resource-manager/Microsoft.DevHub/preview/2022-10-11-preview/examples/Workflow_UpdateTags.json
if __name__ == "__main__":
    main()
