from azure.identity import DefaultAzureCredential
from azure.mgmt.containerservicefleet import ContainerServiceFleetMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-containerservicefleet
# USAGE
    python fleets_list_credentials_result.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ContainerServiceFleetMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid1",
    )

    response = client.fleets.list_credentials(
        resource_group_name="rg1",
        fleet_name="fleet",
    )
    print(response)


# x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/stable/2023-10-15/examples/Fleets_ListCredentialsResult.json
if __name__ == "__main__":
    main()
