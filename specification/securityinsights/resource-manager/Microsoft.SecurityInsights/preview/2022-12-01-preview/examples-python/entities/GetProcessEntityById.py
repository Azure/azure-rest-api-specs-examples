from azure.identity import DefaultAzureCredential
from azure.mgmt.securityinsight import SecurityInsights

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-securityinsight
# USAGE
    python get_process_entity_by_id.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SecurityInsights(
        credential=DefaultAzureCredential(),
        subscription_id="d0cfe6b2-9ac0-4464-9919-dccaee2e48c0",
    )

    response = client.entities.get(
        resource_group_name="myRg",
        workspace_name="myWorkspace",
        entity_id="7264685c-038c-42c6-948c-38e14ef1fb98",
    )
    print(response)


# x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-12-01-preview/examples/entities/GetProcessEntityById.json
if __name__ == "__main__":
    main()
