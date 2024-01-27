
import com.azure.resourcemanager.resources.models.ExportTemplateRequest;
import java.util.Arrays;

/** Samples for ResourceGroups ExportTemplate. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2022-09-01/examples/
     * ExportResourceGroupWithFiltering.json
     */
    /**
     * Sample code: Export a resource group with filtering.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void exportAResourceGroupWithFiltering(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().serviceClient().getResourceGroups().exportTemplate("my-resource-group",
            new ExportTemplateRequest().withResources(Arrays.asList(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/my-resource-group/providers/My.RP/myResourceType/myFirstResource"))
                .withOptions("SkipResourceNameParameterization"),
            com.azure.core.util.Context.NONE);
    }
}
