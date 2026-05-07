from azure.identity import DefaultAzureCredential

from azure.mgmt.web import WebSiteManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-web
# USAGE
    python workflows_regenerate_access_key.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = WebSiteManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    client.workflows.regenerate_access_key(
        resource_group_name="testResourceGroup",
        name="test-name",
        workflow_name="testWorkflowName",
        key_type={"keyType": "Primary"},
    )


# x-ms-original-file: 2025-05-01/Workflows_RegenerateAccessKey.json
if __name__ == "__main__":
    main()
