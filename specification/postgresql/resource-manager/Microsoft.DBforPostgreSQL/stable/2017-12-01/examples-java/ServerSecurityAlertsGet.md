Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-postgresql_1.0.2/sdk/postgresql/azure-resourcemanager-postgresql/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.postgresql.models.SecurityAlertPolicyName;

/** Samples for ServerSecurityAlertPolicies Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/ServerSecurityAlertsGet.json
     */
    /**
     * Sample code: Get a server's threat detection policy.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void getAServerSThreatDetectionPolicy(
        com.azure.resourcemanager.postgresql.PostgreSqlManager manager) {
        manager
            .serverSecurityAlertPolicies()
            .getWithResponse("securityalert-4799", "securityalert-6440", SecurityAlertPolicyName.DEFAULT, Context.NONE);
    }
}
```
