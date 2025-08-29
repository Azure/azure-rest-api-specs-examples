from azure.identity import DefaultAzureCredential

from azure.mgmt.containerservicefleet import ContainerServiceFleetMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-containerservicefleet
# USAGE
    python gates_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ContainerServiceFleetMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.gates.begin_update(
        resource_group_name="rg1",
        fleet_name="fleet1",
        gate_name="12345678-910a-bcde-f000-000000000000",
        properties={"properties": {"state": "Completed"}},
    ).result()
    print(response)


# x-ms-original-file: 2025-04-01-preview/Gates_Update.json
if __name__ == "__main__":
    main()
