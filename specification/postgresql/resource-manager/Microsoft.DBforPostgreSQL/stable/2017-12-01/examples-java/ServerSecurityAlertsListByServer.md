Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-postgresql_1.0.2/sdk/postgresql/azure-resourcemanager-postgresql/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for ServerSecurityAlertPolicies ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/ServerSecurityAlertsListByServer.json
     */
    /**
     * Sample code: List the server's threat detection policies.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void listTheServerSThreatDetectionPolicies(
        com.azure.resourcemanager.postgresql.PostgreSqlManager manager) {
        manager.serverSecurityAlertPolicies().listByServer("securityalert-4799", "securityalert-6440", Context.NONE);
    }
}
```
