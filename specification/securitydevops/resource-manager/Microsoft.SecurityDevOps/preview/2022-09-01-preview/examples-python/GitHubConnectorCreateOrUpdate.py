from azure.identity import DefaultAzureCredential
from azure.mgmt.securitydevops import MicrosoftSecurityDevOps

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-securitydevops
# USAGE
    python git_hub_connector_create_or_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = MicrosoftSecurityDevOps(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.git_hub_connector.begin_create_or_update(
        resource_group_name="westusrg",
        git_hub_connector_name="testconnector",
        git_hub_connector={"location": "West US", "properties": {"code": "00000000000000000000"}},
    ).result()
    print(response)


# x-ms-original-file: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/GitHubConnectorCreateOrUpdate.json
if __name__ == "__main__":
    main()
