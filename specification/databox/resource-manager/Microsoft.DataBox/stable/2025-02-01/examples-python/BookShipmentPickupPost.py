from azure.identity import DefaultAzureCredential

from azure.mgmt.databox import DataBoxManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-databox
# USAGE
    python book_shipment_pickup_post.py

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

    response = client.jobs.book_shipment_pick_up(
        resource_group_name="YourResourceGroupName",
        job_name="TestJobName1",
        shipment_pick_up_request={
            "endTime": "2019-09-22T18:30:00Z",
            "shipmentLocation": "Front desk",
            "startTime": "2019-09-20T18:30:00Z",
        },
    )
    print(response)


# x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/stable/2025-02-01/examples/BookShipmentPickupPost.json
if __name__ == "__main__":
    main()
