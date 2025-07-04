from azure.identity import DefaultAzureCredential

from azure.mgmt.web import WebSiteManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-web
# USAGE
    python diagnostics_get_app_service_certificate_order_detector_response.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = WebSiteManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="5700fc96-77b4-4f8d-afce-c353d8c443bd",
    )

    response = client.certificate_orders_diagnostics.get_app_service_certificate_order_detector_response(
        resource_group_name="Sample-WestUSResourceGroup",
        certificate_order_name="SampleCertificateOrderName",
        detector_name="AutoRenewStatus",
    )
    print(response)


# x-ms-original-file: specification/web/resource-manager/Microsoft.CertificateRegistration/stable/2024-11-01/examples/Diagnostics_GetAppServiceCertificateOrderDetectorResponse.json
if __name__ == "__main__":
    main()
