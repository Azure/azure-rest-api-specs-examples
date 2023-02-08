import com.azure.resourcemanager.providerhub.fluent.models.OperationsDefinitionInner;
import com.azure.resourcemanager.providerhub.models.OperationsDefinitionDisplay;
import com.azure.resourcemanager.providerhub.models.OperationsPutContent;
import java.util.Arrays;

/** Samples for Operations CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/Operations_CreateOrUpdate.json
     */
    /**
     * Sample code: Operations_CreateOrUpdate.
     *
     * @param manager Entry point to ProviderHubManager.
     */
    public static void operationsCreateOrUpdate(com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager
            .operations()
            .createOrUpdateWithResponse(
                "Microsoft.Contoso",
                new OperationsPutContent()
                    .withContents(
                        Arrays
                            .asList(
                                new OperationsDefinitionInner()
                                    .withName("Microsoft.Contoso/Employees/Read")
                                    .withDisplay(
                                        new OperationsDefinitionDisplay()
                                            .withProvider("Microsoft.Contoso")
                                            .withResource("Employees")
                                            .withOperation("Gets/List employee resources")
                                            .withDescription("Read employees")))),
                com.azure.core.util.Context.NONE);
    }
}
