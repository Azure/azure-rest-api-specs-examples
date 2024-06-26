from azure.identity import DefaultAzureCredential
from azure.mgmt.networkanalytics import NetworkAnalyticsMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-networkanalytics
# USAGE
    python data_products_remove_user_role_minimum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = NetworkAnalyticsMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-00000000000",
    )

    client.data_products.remove_user_role(
        resource_group_name="aoiresourceGroupName",
        data_product_name="dataproduct01",
        body={
            "dataTypeScope": ["scope"],
            "principalId": "00000000-0000-0000-0000-00000000000",
            "principalType": "User",
            "role": "Reader",
            "roleAssignmentId": "00000000-0000-0000-0000-00000000000",
            "roleId": "00000000-0000-0000-0000-00000000000",
            "userName": "UserName",
        },
    )


# x-ms-original-file: specification/networkanalytics/resource-manager/Microsoft.NetworkAnalytics/stable/2023-11-15/examples/DataProducts_RemoveUserRole_MinimumSet_Gen.json
if __name__ == "__main__":
    main()
