import com.azure.resourcemanager.networkfunction.models.DestinationType;
import com.azure.resourcemanager.networkfunction.models.EmissionPoliciesPropertiesFormat;
import com.azure.resourcemanager.networkfunction.models.EmissionPolicyDestination;
import com.azure.resourcemanager.networkfunction.models.EmissionType;
import com.azure.resourcemanager.networkfunction.models.IngestionPolicyPropertiesFormat;
import com.azure.resourcemanager.networkfunction.models.IngestionSourcesPropertiesFormat;
import com.azure.resourcemanager.networkfunction.models.IngestionType;
import com.azure.resourcemanager.networkfunction.models.SourceType;
import java.util.Arrays;

/** Samples for CollectorPolicies CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkfunction/resource-manager/Microsoft.NetworkFunction/stable/2022-05-01/examples/CollectorPolicyCreate.json
     */
    /**
     * Sample code: Create a collection policy.
     *
     * @param manager Entry point to AzureTrafficCollectorManager.
     */
    public static void createACollectionPolicy(
        com.azure.resourcemanager.networkfunction.AzureTrafficCollectorManager manager) {
        manager
            .collectorPolicies()
            .define("cp1")
            .withExistingAzureTrafficCollector("rg1", "atc")
            .withIngestionPolicy(
                new IngestionPolicyPropertiesFormat()
                    .withIngestionType(IngestionType.IPFIX)
                    .withIngestionSources(
                        Arrays
                            .asList(
                                new IngestionSourcesPropertiesFormat()
                                    .withSourceType(SourceType.RESOURCE)
                                    .withResourceId(
                                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRouteCircuits/circuitName"))))
            .withEmissionPolicies(
                Arrays
                    .asList(
                        new EmissionPoliciesPropertiesFormat()
                            .withEmissionType(EmissionType.IPFIX)
                            .withEmissionDestinations(
                                Arrays
                                    .asList(
                                        new EmissionPolicyDestination()
                                            .withDestinationType(DestinationType.AZURE_MONITOR)))))
            .create();
    }
}
