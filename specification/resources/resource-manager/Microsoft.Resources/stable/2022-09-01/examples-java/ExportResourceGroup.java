import com.azure.resourcemanager.resources.models.ExportTemplateRequest;
import java.util.Arrays;

/** Samples for ResourceGroups ExportTemplate. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2022-09-01/examples/ExportResourceGroup.json
     */
    /**
     * Sample code: Export a resource group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void exportAResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .genericResources()
            .manager()
            .serviceClient()
            .getResourceGroups()
            .exportTemplate(
                "my-resource-group",
                new ExportTemplateRequest()
                    .withResources(Arrays.asList("*"))
                    .withOptions("IncludeParameterDefaultValue,IncludeComments"),
                com.azure.core.util.Context.NONE);
    }
}
