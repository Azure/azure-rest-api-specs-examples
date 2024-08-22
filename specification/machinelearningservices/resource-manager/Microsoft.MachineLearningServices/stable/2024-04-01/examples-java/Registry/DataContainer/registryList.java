
import com.azure.resourcemanager.machinelearning.models.ListViewType;

/**
 * Samples for RegistryDataContainers List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Registry/DataContainer/registryList.json
     */
    /**
     * Sample code: RegistryList Registry Data Container.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        registryListRegistryDataContainer(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.registryDataContainers().list("test-rg", "registryName", null, ListViewType.ALL,
            com.azure.core.util.Context.NONE);
    }
}
