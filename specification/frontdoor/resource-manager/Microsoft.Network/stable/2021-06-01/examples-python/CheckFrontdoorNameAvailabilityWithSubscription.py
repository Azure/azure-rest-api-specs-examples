from azure.identity import DefaultAzureCredential
from azure.mgmt.frontdoor import FrontDoorManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-frontdoor
# USAGE
    python check_frontdoor_name_availability_with_subscription.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = FrontDoorManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    response = client.front_door_name_availability_with_subscription.check(
        check_front_door_name_availability_input={
            "name": "sampleName",
            "type": "Microsoft.Network/frontDoors/frontendEndpoints",
        },
    )
    print(response)


# x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2021-06-01/examples/CheckFrontdoorNameAvailabilityWithSubscription.json
if __name__ == "__main__":
    main()
