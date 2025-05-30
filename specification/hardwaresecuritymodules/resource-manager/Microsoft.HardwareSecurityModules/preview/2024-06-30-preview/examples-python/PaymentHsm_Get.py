from azure.identity import DefaultAzureCredential

from azure.mgmt.hardwaresecuritymodules import HardwareSecurityModulesMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-hardwaresecuritymodules
# USAGE
    python payment_hsm_get.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = HardwareSecurityModulesMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.dedicated_hsm.get(
        resource_group_name="hsm-group",
        name="hsm1",
    )
    print(response)


# x-ms-original-file: specification/hardwaresecuritymodules/resource-manager/Microsoft.HardwareSecurityModules/preview/2024-06-30-preview/examples/PaymentHsm_Get.json
if __name__ == "__main__":
    main()
