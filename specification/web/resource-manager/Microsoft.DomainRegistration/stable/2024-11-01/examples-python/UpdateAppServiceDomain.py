from azure.identity import DefaultAzureCredential

from azure.mgmt.web import WebSiteManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-web
# USAGE
    python update_app_service_domain.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = WebSiteManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="34adfa4f-cedf-4dc0-ba29-b6d1a69ab345",
    )

    response = client.domains.update(
        resource_group_name="testrg123",
        domain_name="example.com",
        domain={
            "properties": {
                "authCode": "exampleAuthCode",
                "autoRenew": True,
                "consent": {
                    "agreedAt": "2021-09-10T19:30:53Z",
                    "agreedBy": "192.0.2.1",
                    "agreementKeys": ["agreementKey1"],
                },
                "contactAdmin": {
                    "addressMailing": {
                        "address1": "3400 State St",
                        "city": "Chicago",
                        "country": "United States",
                        "postalCode": "67098",
                        "state": "IL",
                    },
                    "email": "admin@email.com",
                    "fax": "1-245-534-2242",
                    "jobTitle": "Admin",
                    "nameFirst": "John",
                    "nameLast": "Doe",
                    "nameMiddle": "",
                    "organization": "Microsoft Inc.",
                    "phone": "1-245-534-2242",
                },
                "contactBilling": {
                    "addressMailing": {
                        "address1": "3400 State St",
                        "city": "Chicago",
                        "country": "United States",
                        "postalCode": "67098",
                        "state": "IL",
                    },
                    "email": "billing@email.com",
                    "fax": "1-245-534-2242",
                    "jobTitle": "Billing",
                    "nameFirst": "John",
                    "nameLast": "Doe",
                    "nameMiddle": "",
                    "organization": "Microsoft Inc.",
                    "phone": "1-245-534-2242",
                },
                "contactRegistrant": {
                    "addressMailing": {
                        "address1": "3400 State St",
                        "city": "Chicago",
                        "country": "United States",
                        "postalCode": "67098",
                        "state": "IL",
                    },
                    "email": "registrant@email.com",
                    "fax": "1-245-534-2242",
                    "jobTitle": "Registrant",
                    "nameFirst": "John",
                    "nameLast": "Doe",
                    "nameMiddle": "",
                    "organization": "Microsoft Inc.",
                    "phone": "1-245-534-2242",
                },
                "contactTech": {
                    "addressMailing": {
                        "address1": "3400 State St",
                        "city": "Chicago",
                        "country": "United States",
                        "postalCode": "67098",
                        "state": "IL",
                    },
                    "email": "tech@email.com",
                    "fax": "1-245-534-2242",
                    "jobTitle": "Tech",
                    "nameFirst": "John",
                    "nameLast": "Doe",
                    "nameMiddle": "",
                    "organization": "Microsoft Inc.",
                    "phone": "1-245-534-2242",
                },
                "dnsType": "DefaultDomainRegistrarDns",
                "privacy": False,
            }
        },
    )
    print(response)


# x-ms-original-file: specification/web/resource-manager/Microsoft.DomainRegistration/stable/2024-11-01/examples/UpdateAppServiceDomain.json
if __name__ == "__main__":
    main()
