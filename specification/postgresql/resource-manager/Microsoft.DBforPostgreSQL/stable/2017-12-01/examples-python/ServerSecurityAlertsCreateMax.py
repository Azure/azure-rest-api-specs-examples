from typing import Any, IO, TYPE_CHECKING, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.rdbms.postgresql import PostgreSQLManagementClient

if TYPE_CHECKING:
    # pylint: disable=unused-import,ungrouped-imports
    from .. import models as _models
"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-rdbms
# USAGE
    python server_security_alerts_create_max.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = PostgreSQLManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-1111-2222-3333-444444444444",
    )

    response = client.server_security_alert_policies.begin_create_or_update(
        resource_group_name="securityalert-4799",
        server_name="securityalert-6440",
        security_alert_policy_name="Default",
        parameters={
            "properties": {
                "disabledAlerts": ["Access_Anomaly", "Usage_Anomaly"],
                "emailAccountAdmins": True,
                "emailAddresses": ["testSecurityAlert@microsoft.com"],
                "retentionDays": 5,
                "state": "Enabled",
                "storageAccountAccessKey": "sdlfkjabc+sdlfkjsdlkfsjdfLDKFTERLKFDFKLjsdfksjdflsdkfD2342309432849328476458/3RSD==",
                "storageEndpoint": "https://mystorage.blob.core.windows.net",
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/ServerSecurityAlertsCreateMax.json
if __name__ == "__main__":
    main()
