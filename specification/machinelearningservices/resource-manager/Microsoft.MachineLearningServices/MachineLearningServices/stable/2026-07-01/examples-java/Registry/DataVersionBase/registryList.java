
import com.azure.resourcemanager.machinelearning.models.ListViewType;

/**
 * Samples for RegistryDataVersions List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Registry/DataVersionBase/registryList.json
     */
    /**
     * Sample code: RegistryList Registry Data Version Base.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        registryListRegistryDataVersionBase(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.registryDataVersions().list("test-rg", "registryName", "string", "string", 1, null, "string",
            ListViewType.ARCHIVED_ONLY, com.azure.core.util.Context.NONE);
    }
}
