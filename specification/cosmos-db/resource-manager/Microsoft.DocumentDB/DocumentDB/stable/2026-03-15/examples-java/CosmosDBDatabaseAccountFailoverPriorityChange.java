
import com.azure.resourcemanager.cosmos.models.FailoverPolicies;
import com.azure.resourcemanager.cosmos.models.FailoverPolicy;
import java.util.Arrays;

/**
 * Samples for DatabaseAccounts FailoverPriorityChange.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBDatabaseAccountFailoverPriorityChange.json
     */
    /**
     * Sample code: CosmosDBDatabaseAccountFailoverPriorityChange.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void
        cosmosDBDatabaseAccountFailoverPriorityChange(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getDatabaseAccounts().failoverPriorityChange("rg1", "ddb1-failover",
            new FailoverPolicies().withFailoverPolicies(
                Arrays.asList(new FailoverPolicy().withLocationName("eastus").withFailoverPriority(0),
                    new FailoverPolicy().withLocationName("westus").withFailoverPriority(1))),
            com.azure.core.util.Context.NONE);
    }
}
