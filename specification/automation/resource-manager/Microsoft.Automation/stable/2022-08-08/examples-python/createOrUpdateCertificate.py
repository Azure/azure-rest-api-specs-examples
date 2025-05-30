from azure.identity import DefaultAzureCredential

from azure.mgmt.automation import AutomationClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-automation
# USAGE
    python create_or_update_certificate.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AutomationClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    response = client.certificate.create_or_update(
        resource_group_name="rg",
        automation_account_name="myAutomationAccount18",
        certificate_name="testCert",
        parameters={
            "name": "testCert",
            "properties": {
                "base64Value": "base 64 value of cert",
                "description": "Sample Cert",
                "isExportable": False,
                "thumbprint": "thumbprint of cert",
            },
        },
    )
    print(response)


# x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2022-08-08/examples/createOrUpdateCertificate.json
if __name__ == "__main__":
    main()
