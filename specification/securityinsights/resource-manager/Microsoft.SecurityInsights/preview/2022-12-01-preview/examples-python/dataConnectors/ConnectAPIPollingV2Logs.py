from azure.identity import DefaultAzureCredential
from azure.mgmt.securityinsight import SecurityInsights

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-securityinsight
# USAGE
    python connect_api_polling_v2_logs.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SecurityInsights(
        credential=DefaultAzureCredential(),
        subscription_id="d0cfe6b2-9ac0-4464-9919-dccaee2e48c0",
    )

    response = client.data_connectors.connect(
        resource_group_name="myRg",
        workspace_name="myWorkspace",
        data_connector_id="316ec55e-7138-4d63-ab18-90c8a60fd1c8",
        connect_body={
            "apiKey": "123456789",
            "dataCollectionEndpoint": "https://test.eastus.ingest.monitor.azure.com",
            "dataCollectionRuleImmutableId": "dcr-34adsj9o7d6f9de204478b9cgb43b631",
            "kind": "APIKey",
            "outputStream": "Custom-MyTableRawData",
            "requestConfigUserInputValues": [
                {
                    "displayText": "Organization Name",
                    "placeHolderName": "{{placeHolder1}}",
                    "placeHolderValue": "somePlaceHolderValue",
                    "requestObjectKey": "apiEndpoint",
                }
            ],
        },
    )
    print(response)


# x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-12-01-preview/examples/dataConnectors/ConnectAPIPollingV2Logs.json
if __name__ == "__main__":
    main()
