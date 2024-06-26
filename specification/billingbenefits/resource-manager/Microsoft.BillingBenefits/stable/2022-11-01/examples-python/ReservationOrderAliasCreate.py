from azure.identity import DefaultAzureCredential
from azure.mgmt.billingbenefits import BillingBenefitsRP

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-billingbenefits
# USAGE
    python reservation_order_alias_create.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = BillingBenefitsRP(
        credential=DefaultAzureCredential(),
    )

    response = client.reservation_order_alias.begin_create(
        reservation_order_alias_name="reservationOrderAlias123",
        body={
            "location": "eastus",
            "properties": {
                "appliedScopeProperties": {
                    "resourceGroupId": "/subscriptions/10000000-0000-0000-0000-000000000000/resourceGroups/testrg"
                },
                "appliedScopeType": "Single",
                "billingPlan": "P1M",
                "billingScopeId": "/subscriptions/10000000-0000-0000-0000-000000000000",
                "displayName": "ReservationOrder_2022-06-02",
                "quantity": 5,
                "renew": True,
                "reservedResourceProperties": {"instanceFlexibility": "On"},
                "reservedResourceType": "VirtualMachines",
                "term": "P1Y",
            },
            "sku": {"name": "Standard_M64s_v2"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/billingbenefits/resource-manager/Microsoft.BillingBenefits/stable/2022-11-01/examples/ReservationOrderAliasCreate.json
if __name__ == "__main__":
    main()
