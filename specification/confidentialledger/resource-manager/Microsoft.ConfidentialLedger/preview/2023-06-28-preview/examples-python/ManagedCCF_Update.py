from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.confidentialledger import ConfidentialLedger

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-confidentialledger
# USAGE
    python managed_ccf_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ConfidentialLedger(
        credential=DefaultAzureCredential(),
        subscription_id="0000000-0000-0000-0000-000000000001",
    )

    response = client.managed_ccf.begin_update(
        resource_group_name="DummyResourceGroupName",
        app_name="DummyMccfAppName",
        managed_ccf={
            "location": "EastUS",
            "properties": {
                "deploymentType": {
                    "appSourceUri": "https://myaccount.blob.core.windows.net/storage/mccfsource?sv=2022-02-11%st=2022-03-11",
                    "languageRuntime": "CPP",
                }
            },
            "tags": {"additionalProps1": "additional properties"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/confidentialledger/resource-manager/Microsoft.ConfidentialLedger/preview/2023-06-28-preview/examples/ManagedCCF_Update.json
if __name__ == "__main__":
    main()
