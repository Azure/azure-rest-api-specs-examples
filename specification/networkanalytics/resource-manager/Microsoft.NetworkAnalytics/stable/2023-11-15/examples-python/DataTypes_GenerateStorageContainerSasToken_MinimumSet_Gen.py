from azure.identity import DefaultAzureCredential
from azure.mgmt.networkanalytics import NetworkAnalyticsMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-networkanalytics
# USAGE
    python data_types_generate_storage_container_sas_token_minimum_set_gen.py

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

    response = client.data_types.generate_storage_container_sas_token(
        resource_group_name="aoiresourceGroupName",
        data_product_name="dataproduct01",
        data_type_name="datatypename",
        body={
            "expiryTimeStamp": "2023-08-24T05:35:16.887Z",
            "ipAddress": "1.1.1.1",
            "startTimeStamp": "2023-08-24T05:35:16.887Z",
        },
    )
    print(response)


# x-ms-original-file: specification/networkanalytics/resource-manager/Microsoft.NetworkAnalytics/stable/2023-11-15/examples/DataTypes_GenerateStorageContainerSasToken_MinimumSet_Gen.json
if __name__ == "__main__":
    main()
