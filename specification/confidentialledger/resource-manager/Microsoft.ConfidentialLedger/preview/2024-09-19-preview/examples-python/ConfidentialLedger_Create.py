from azure.identity import DefaultAzureCredential

from azure.mgmt.confidentialledger import ConfidentialLedger

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-confidentialledger
# USAGE
    python confidential_ledger_create.py

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

    response = client.ledger.begin_create(
        resource_group_name="DummyResourceGroupName",
        ledger_name="DummyLedgerName",
        confidential_ledger={
            "location": "EastUS",
            "properties": {
                "aadBasedSecurityPrincipals": [
                    {
                        "ledgerRoleName": "Administrator",
                        "principalId": "34621747-6fc8-4771-a2eb-72f31c461f2e",
                        "tenantId": "bce123b9-2b7b-4975-8360-5ca0b9b1cd08",
                    }
                ],
                "certBasedSecurityPrincipals": [
                    {
                        "cert": "-----BEGIN CERTIFICATE-----MIIBsjCCATigAwIBAgIUZWIbyG79TniQLd2UxJuU74tqrKcwCgYIKoZIzj0EAwMwEDEOMAwGA1UEAwwFdXNlcjAwHhcNMjEwMzE2MTgwNjExWhcNMjIwMzE2MTgwNjExWjAQMQ4wDAYDVQQDDAV1c2VyMDB2MBAGByqGSM49AgEGBSuBBAAiA2IABBiWSo/j8EFit7aUMm5lF+lUmCu+IgfnpFD+7QMgLKtxRJ3aGSqgS/GpqcYVGddnODtSarNE/HyGKUFUolLPQ5ybHcouUk0kyfA7XMeSoUA4lBz63Wha8wmXo+NdBRo39qNTMFEwHQYDVR0OBBYEFPtuhrwgGjDFHeUUT4nGsXaZn69KMB8GA1UdIwQYMBaAFPtuhrwgGjDFHeUUT4nGsXaZn69KMA8GA1UdEwEB/wQFMAMBAf8wCgYIKoZIzj0EAwMDaAAwZQIxAOnozm2CyqRwSSQLls5r+mUHRGRyXHXwYtM4Dcst/VEZdmS9fqvHRCHbjUlO/+HNfgIwMWZ4FmsjD3wnPxONOm9YdVn/PRD7SsPRPbOjwBiE4EBGaHDsLjYAGDSGi7NJnSkA-----END CERTIFICATE-----",
                        "ledgerRoleName": "Reader",
                    }
                ],
                "hostLevel": "Info",
                "ledgerSku": "Standard",
                "ledgerType": "Public",
                "maxBodySizeInMb": 1,
                "nodeCount": 3,
                "subjectName": "CN=CCF Node",
                "workerThreads": 0,
                "writeLBAddressPrefix": "write",
            },
            "tags": {"additionalProps1": "additional properties"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/confidentialledger/resource-manager/Microsoft.ConfidentialLedger/preview/2024-09-19-preview/examples/ConfidentialLedger_Create.json
if __name__ == "__main__":
    main()
