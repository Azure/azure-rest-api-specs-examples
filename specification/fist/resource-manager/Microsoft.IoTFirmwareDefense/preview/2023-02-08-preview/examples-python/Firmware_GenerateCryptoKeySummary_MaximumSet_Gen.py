from azure.identity import DefaultAzureCredential
from azure.mgmt.iotfirmwaredefense import IoTFirmwareDefenseMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-iotfirmwaredefense
# USAGE
    python firmware_generate_crypto_key_summary_maximum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = IoTFirmwareDefenseMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="C589E84A-5C11-4A25-9CF9-4E9C2F1EBFCA",
    )

    response = client.firmware.generate_crypto_key_summary(
        resource_group_name="FirmwareAnalysisRG",
        workspace_name="default",
        firmware_id="DECAFBAD-0000-0000-0000-BADBADBADBAD",
    )
    print(response)


# x-ms-original-file: specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/preview/2023-02-08-preview/examples/Firmware_GenerateCryptoKeySummary_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
