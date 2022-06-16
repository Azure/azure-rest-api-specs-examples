import com.azure.core.util.Context;
import com.azure.resourcemanager.synapse.fluent.models.LanguageExtensionInner;
import com.azure.resourcemanager.synapse.models.LanguageExtensionName;
import com.azure.resourcemanager.synapse.models.LanguageExtensionsList;
import java.util.Arrays;

/** Samples for KustoPools AddLanguageExtensions. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolLanguageExtensionsAdd.json
     */
    /**
     * Sample code: KustoPoolAddLanguageExtensions.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void kustoPoolAddLanguageExtensions(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .kustoPools()
            .addLanguageExtensions(
                "kustorptest",
                "kustoclusterrptest4",
                "kustorptest",
                new LanguageExtensionsList()
                    .withValue(
                        Arrays
                            .asList(
                                new LanguageExtensionInner().withLanguageExtensionName(LanguageExtensionName.PYTHON),
                                new LanguageExtensionInner().withLanguageExtensionName(LanguageExtensionName.R))),
                Context.NONE);
    }
}
