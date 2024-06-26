from azure.identity import DefaultAzureCredential
from azure.mgmt.reservations import AzureReservationAPI

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-reservations
# USAGE
    python patch_compute_quota_request.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AzureReservationAPI(
        credential=DefaultAzureCredential(),
    )

    response = client.quota.begin_update(
        subscription_id="D7EC67B3-7657-4966-BFFC-41EFD36BAAB3",
        provider_id="Microsoft.Compute",
        location="eastus",
        resource_name="standardFSv2Family",
        create_quota_request={"properties": {"limit": 200, "name": {"value": "standardFSv2Family"}, "unit": "Count"}},
    ).result()
    print(response)


# x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2020-10-25/examples/patchComputeQuotaRequest.json
if __name__ == "__main__":
    main()
