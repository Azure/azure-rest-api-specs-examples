from azure.identity import DefaultAzureCredential

from azure.mgmt.hardwaresecuritymodules import HardwareSecurityModulesMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-hardwaresecuritymodules
# USAGE
    python dedicated_hsm_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = HardwareSecurityModulesMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.dedicated_hsm.begin_update(
        resource_group_name="hsm-group",
        name="hsm1",
        parameters={"tags": {"Dept": "hsm", "Environment": "dogfood", "Slice": "A"}},
    ).result()
    print(response)


# x-ms-original-file: 2025-03-31/DedicatedHsm_Update.json
if __name__ == "__main__":
    main()
