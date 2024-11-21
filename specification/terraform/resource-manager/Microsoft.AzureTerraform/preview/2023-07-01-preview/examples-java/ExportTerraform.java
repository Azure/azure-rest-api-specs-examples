
import com.azure.resourcemanager.terraform.models.ExportResourceGroup;

/**
 * Samples for Terraform ExportTerraform.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-07-01-preview/ExportTerraform.json
     */
    /**
     * Sample code: ExportTerraform.
     * 
     * @param manager Entry point to AzureTerraformManager.
     */
    public static void exportTerraform(com.azure.resourcemanager.terraform.AzureTerraformManager manager) {
        manager.terraforms().exportTerraform(new ExportResourceGroup().withResourceGroupName("rg1"),
            com.azure.core.util.Context.NONE);
    }
}
