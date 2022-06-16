import com.azure.core.util.Context;

/** Samples for KustoPools ListLanguageExtensions. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolLanguageExtensionsList.json
     */
    /**
     * Sample code: KustoPoolListLanguageExtensions.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void kustoPoolListLanguageExtensions(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.kustoPools().listLanguageExtensions("kustorptest", "kustoclusterrptest4", "kustorptest", Context.NONE);
    }
}
