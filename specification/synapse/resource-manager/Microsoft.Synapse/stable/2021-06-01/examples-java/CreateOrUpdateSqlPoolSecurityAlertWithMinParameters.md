Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.5/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.synapse.models.SecurityAlertPolicyName;
import com.azure.resourcemanager.synapse.models.SecurityAlertPolicyState;
import com.azure.resourcemanager.synapse.models.SqlPoolSecurityAlertPolicy;

/** Samples for SqlPoolSecurityAlertPolicies CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateOrUpdateSqlPoolSecurityAlertWithMinParameters.json
     */
    /**
     * Sample code: Update a Sql pool's threat detection policy with minimal parameters.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void updateASqlPoolSThreatDetectionPolicyWithMinimalParameters(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        SqlPoolSecurityAlertPolicy resource =
            manager
                .sqlPoolSecurityAlertPolicies()
                .getWithResponse(
                    "securityalert-4799", "securityalert-6440", "testdb", SecurityAlertPolicyName.DEFAULT, Context.NONE)
                .getValue();
        resource.update().withState(SecurityAlertPolicyState.ENABLED).apply();
    }
}
```
