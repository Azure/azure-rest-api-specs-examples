from azure.identity import DefaultAzureCredential
from azure.mgmt.deviceupdate import DeviceUpdateMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-deviceupdate
# USAGE
    python check_name_availability_already_exists.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DeviceUpdateMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.check_name_availability(
        request={"name": "contoso", "type": "Microsoft.DeviceUpdate/accounts"},
    )
    print(response)


# x-ms-original-file: specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/stable/2022-10-01/examples/CheckNameAvailability_AlreadyExists.json
if __name__ == "__main__":
    main()
