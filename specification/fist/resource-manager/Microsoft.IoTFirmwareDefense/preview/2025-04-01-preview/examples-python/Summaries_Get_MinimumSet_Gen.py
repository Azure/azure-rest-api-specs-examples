from azure.identity import DefaultAzureCredential

from azure.mgmt.iotfirmwaredefense import IoTFirmwareDefenseMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-iotfirmwaredefense
# USAGE
    python summaries_get_minimum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = IoTFirmwareDefenseMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.summaries.get(
        resource_group_name="FirmwareAnalysisRG",
        workspace_name="default",
        firmware_id="109a9886-50bf-85a8-9d75-000000000000",
        summary_type="Firmware",
    )
    print(response)


# x-ms-original-file: 2025-04-01-preview/Summaries_Get_MinimumSet_Gen.json
if __name__ == "__main__":
    main()
