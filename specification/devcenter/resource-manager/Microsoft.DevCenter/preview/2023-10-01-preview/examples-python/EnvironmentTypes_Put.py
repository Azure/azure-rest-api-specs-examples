from azure.identity import DefaultAzureCredential
from azure.mgmt.devcenter import DevCenterMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-devcenter
# USAGE
    python environment_types_put.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DevCenterMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="0ac520ee-14c0-480f-b6c9-0a90c58ffff",
    )

    response = client.environment_types.create_or_update(
        resource_group_name="rg1",
        dev_center_name="Contoso",
        environment_type_name="DevTest",
        body={"properties": {"displayName": "Dev"}, "tags": {"Owner": "superuser"}},
    )
    print(response)


# x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2023-10-01-preview/examples/EnvironmentTypes_Put.json
if __name__ == "__main__":
    main()
