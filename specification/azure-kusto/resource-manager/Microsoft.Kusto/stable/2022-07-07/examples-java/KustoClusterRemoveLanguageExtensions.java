import com.azure.core.util.Context;
import com.azure.resourcemanager.kusto.fluent.models.LanguageExtensionInner;
import com.azure.resourcemanager.kusto.models.LanguageExtensionName;
import com.azure.resourcemanager.kusto.models.LanguageExtensionsList;
import java.util.Arrays;

/** Samples for Clusters RemoveLanguageExtensions. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-07-07/examples/KustoClusterRemoveLanguageExtensions.json
     */
    /**
     * Sample code: KustoClusterRemoveLanguageExtensions.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoClusterRemoveLanguageExtensions(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .clusters()
            .removeLanguageExtensions(
                "kustorptest",
                "kustoCluster",
                new LanguageExtensionsList()
                    .withValue(
                        Arrays
                            .asList(
                                new LanguageExtensionInner().withLanguageExtensionName(LanguageExtensionName.PYTHON),
                                new LanguageExtensionInner().withLanguageExtensionName(LanguageExtensionName.R))),
                Context.NONE);
    }
}
