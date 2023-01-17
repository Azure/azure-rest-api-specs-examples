from azure.identity import DefaultAzureCredential
from azure.mgmt.dataprotection import DataProtectionClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-dataprotection
# USAGE
    python checkfeature_support.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DataProtectionClient(
        credential=DefaultAzureCredential(),
        subscription_id="0b352192-dcac-4cc7-992e-a96190ccc68c",
    )

    response = client.data_protection.check_feature_support(
        location="WestUS",
        parameters={"featureType": "DataSourceType", "objectType": "FeatureValidationRequest"},
    )
    print(response)


# x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-11-01-preview/examples/CheckfeatureSupport.json
if __name__ == "__main__":
    main()
