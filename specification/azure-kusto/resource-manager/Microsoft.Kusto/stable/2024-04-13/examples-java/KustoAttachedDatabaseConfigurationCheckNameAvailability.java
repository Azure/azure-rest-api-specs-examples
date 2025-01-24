
import com.azure.resourcemanager.kusto.models.AttachedDatabaseConfigurationsCheckNameRequest;

/**
 * Samples for AttachedDatabaseConfigurations CheckNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2024-04-13/examples/
     * KustoAttachedDatabaseConfigurationCheckNameAvailability.json
     */
    /**
     * Sample code: KustoAttachedDatabaseConfigurationCheckNameAvailability.
     * 
     * @param manager Entry point to KustoManager.
     */
    public static void
        kustoAttachedDatabaseConfigurationCheckNameAvailability(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.attachedDatabaseConfigurations().checkNameAvailabilityWithResponse("kustorptest", "kustoCluster",
            new AttachedDatabaseConfigurationsCheckNameRequest().withName("adc1"), com.azure.core.util.Context.NONE);
    }
}
