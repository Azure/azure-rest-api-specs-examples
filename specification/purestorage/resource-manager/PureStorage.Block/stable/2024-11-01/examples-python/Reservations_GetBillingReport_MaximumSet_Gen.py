from azure.identity import DefaultAzureCredential

from azure.mgmt.purestorageblock import PureStorageBlockMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-purestorageblock
# USAGE
    python reservations_get_billing_report_maximum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = PureStorageBlockMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.reservations.get_billing_report(
        resource_group_name="rgpurestorage",
        reservation_name="reservationname",
    )
    print(response)


# x-ms-original-file: 2024-11-01/Reservations_GetBillingReport_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
