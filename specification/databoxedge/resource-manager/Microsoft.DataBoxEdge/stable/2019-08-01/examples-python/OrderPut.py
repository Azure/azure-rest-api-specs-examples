from azure.identity import DefaultAzureCredential

from azure.mgmt.databoxedge import DataBoxEdgeManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-databoxedge
# USAGE
    python order_put.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DataBoxEdgeManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="4385cf00-2d3a-425a-832f-f4285b1c9dce",
    )

    response = client.orders.begin_create_or_update(
        device_name="testedgedevice",
        resource_group_name="GroupForEdgeAutomation",
        order={
            "properties": {
                "contactInformation": {
                    "companyName": "Microsoft",
                    "contactPerson": "John Mcclane",
                    "emailList": ["john@microsoft.com"],
                    "phone": "(800) 426-9400",
                },
                "shippingAddress": {
                    "addressLine1": "Microsoft Corporation",
                    "addressLine2": "One Microsoft Way",
                    "addressLine3": "Redmond",
                    "city": "WA",
                    "country": "USA",
                    "postalCode": "98052",
                    "state": "WA",
                },
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/OrderPut.json
if __name__ == "__main__":
    main()
