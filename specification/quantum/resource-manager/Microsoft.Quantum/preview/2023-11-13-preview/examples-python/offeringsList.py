from azure.identity import DefaultAzureCredential
from azure.mgmt.quantum import AzureQuantumMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-quantum
# USAGE
    python offerings_list.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AzureQuantumMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="1C4B2828-7D49-494F-933D-061373BE28C2",
    )

    response = client.offerings.list(
        location_name="westus2",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/quantum/resource-manager/Microsoft.Quantum/preview/2023-11-13-preview/examples/offeringsList.json
if __name__ == "__main__":
    main()
