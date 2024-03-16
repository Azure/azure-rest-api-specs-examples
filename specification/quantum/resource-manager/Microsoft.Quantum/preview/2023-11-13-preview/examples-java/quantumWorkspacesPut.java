
import com.azure.resourcemanager.quantum.models.Provider;
import java.util.Arrays;

/**
 * Samples for Workspaces CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/quantum/resource-manager/Microsoft.Quantum/preview/2023-11-13-preview/examples/quantumWorkspacesPut
     * .json
     */
    /**
     * Sample code: QuantumWorkspacesPut.
     * 
     * @param manager Entry point to AzureQuantumManager.
     */
    public static void quantumWorkspacesPut(com.azure.resourcemanager.quantum.AzureQuantumManager manager) {
        manager.workspaces().define("quantumworkspace1").withRegion("West US")
            .withExistingResourceGroup("quantumResourcegroup")
            .withProviders(Arrays.asList(new Provider().withProviderId("Honeywell").withProviderSku("Basic"),
                new Provider().withProviderId("IonQ").withProviderSku("Basic"),
                new Provider().withProviderId("OneQBit").withProviderSku("Basic")))
            .withStorageAccount(
                "/subscriptions/1C4B2828-7D49-494F-933D-061373BE28C2/resourceGroups/quantumResourcegroup/providers/Microsoft.Storage/storageAccounts/testStorageAccount")
            .create();
    }
}
