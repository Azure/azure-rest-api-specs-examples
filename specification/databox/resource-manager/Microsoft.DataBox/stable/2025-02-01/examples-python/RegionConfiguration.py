from azure.identity import DefaultAzureCredential

from azure.mgmt.databox import DataBoxManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-databox
# USAGE
    python region_configuration.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DataBoxManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="YourSubscriptionId",
    )

    response = client.service.region_configuration(
        location="westus",
        region_configuration_request={
            "deviceCapabilityRequest": {"model": "DataBoxDisk", "skuName": "DataBoxDisk"},
            "scheduleAvailabilityRequest": {"model": "DataBox", "skuName": "DataBox", "storageLocation": "westus"},
        },
    )
    print(response)


# x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/stable/2025-02-01/examples/RegionConfiguration.json
if __name__ == "__main__":
    main()
