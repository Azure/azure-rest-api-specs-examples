from azure.identity import DefaultAzureCredential
from azure.mgmt.peering import PeeringManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-peering
# USAGE
    python call_looking_glass_to_execute_a_command.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = PeeringManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subId",
    )

    response = client.looking_glass.invoke(
        command="Traceroute",
        source_type="AzureRegion",
        source_location="West US",
        destination_ip="0.0.0.0",
    )
    print(response)


# x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2022-10-01/examples/LookingGlassInvokeCommand.json
if __name__ == "__main__":
    main()
