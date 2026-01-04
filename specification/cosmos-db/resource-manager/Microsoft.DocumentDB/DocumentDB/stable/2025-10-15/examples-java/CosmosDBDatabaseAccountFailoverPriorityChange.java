
import com.azure.resourcemanager.cosmos.models.FailoverPolicies;
import com.azure.resourcemanager.cosmos.models.FailoverPolicy;
import java.util.Arrays;

/**
 * Samples for DatabaseAccounts FailoverPriorityChange.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBDatabaseAccountFailoverPriorityChange.json
     */
    /**
     * Sample code: CosmosDBDatabaseAccountFailoverPriorityChange.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        cosmosDBDatabaseAccountFailoverPriorityChange(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getDatabaseAccounts().failoverPriorityChange("rg1",
            "ddb1-failover",
            new FailoverPolicies().withFailoverPolicies(
                Arrays.asList(new FailoverPolicy().withLocationName("eastus").withFailoverPriority(0),
                    new FailoverPolicy().withLocationName("westus").withFailoverPriority(1))),
            com.azure.core.util.Context.NONE);
    }
}
