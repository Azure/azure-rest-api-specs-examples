```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.synapse.models.SecurityAlertPolicyName;
import com.azure.resourcemanager.synapse.models.SecurityAlertPolicyState;
import com.azure.resourcemanager.synapse.models.SqlPoolSecurityAlertPolicy;
import java.util.Arrays;

/** Samples for SqlPoolSecurityAlertPolicies CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateOrUpdateSqlPoolSecurityAlertWithAllParameters.json
     */
    /**
     * Sample code: Update a Sql pool's threat detection policy with all parameters.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void updateASqlPoolSThreatDetectionPolicyWithAllParameters(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        SqlPoolSecurityAlertPolicy resource =
            manager
                .sqlPoolSecurityAlertPolicies()
                .getWithResponse(
                    "securityalert-4799", "securityalert-6440", "testdb", SecurityAlertPolicyName.DEFAULT, Context.NONE)
                .getValue();
        resource
            .update()
            .withState(SecurityAlertPolicyState.ENABLED)
            .withDisabledAlerts(Arrays.asList("Sql_Injection", "Usage_Anomaly"))
            .withEmailAddresses(Arrays.asList("test@microsoft.com", "user@microsoft.com"))
            .withEmailAccountAdmins(true)
            .withStorageEndpoint("https://mystorage.blob.core.windows.net")
            .withStorageAccountAccessKey(
                "sdlfkjabc+sdlfkjsdlkfsjdfLDKFTERLKFDFKLjsdfksjdflsdkfD2342309432849328476458/3RSD==")
            .withRetentionDays(6)
            .apply();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.6/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.
