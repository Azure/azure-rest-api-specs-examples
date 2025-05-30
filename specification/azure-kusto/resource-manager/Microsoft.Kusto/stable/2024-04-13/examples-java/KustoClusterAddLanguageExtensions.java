
import com.azure.resourcemanager.kusto.fluent.models.LanguageExtensionInner;
import com.azure.resourcemanager.kusto.models.LanguageExtensionName;
import com.azure.resourcemanager.kusto.models.LanguageExtensionsList;
import java.util.Arrays;

/**
 * Samples for Clusters AddLanguageExtensions.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2024-04-13/examples/
     * KustoClusterAddLanguageExtensions.json
     */
    /**
     * Sample code: KustoClusterAddLanguageExtensions.
     * 
     * @param manager Entry point to KustoManager.
     */
    public static void kustoClusterAddLanguageExtensions(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.clusters().addLanguageExtensions("kustorptest", "kustoCluster",
            new LanguageExtensionsList().withValue(
                Arrays.asList(new LanguageExtensionInner().withLanguageExtensionName(LanguageExtensionName.PYTHON),
                    new LanguageExtensionInner().withLanguageExtensionName(LanguageExtensionName.R))),
            com.azure.core.util.Context.NONE);
    }
}
