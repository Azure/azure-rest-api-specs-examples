
import com.azure.resourcemanager.apimanagement.models.KeyVaultContractCreateProperties;
import com.azure.resourcemanager.apimanagement.models.NamedValueCreateContract;
import java.util.Arrays;

/**
 * Samples for WorkspaceNamedValue CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementCreateWorkspaceNamedValueWithKeyVault.json
     */
    /**
     * Sample code: ApiManagementCreateWorkspaceNamedValueWithKeyVault.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateWorkspaceNamedValueWithKeyVault(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceNamedValues().createOrUpdate("rg1", "apimService1", "wks1", "testprop6",
            new NamedValueCreateContract().withDisplayName("prop6namekv")
                .withKeyVault(new KeyVaultContractCreateProperties().withSecretIdentifier("fakeTokenPlaceholder")
                    .withIdentityClientId("ceaa6b06-c00f-43ef-99ac-f53d1fe876a0"))
                .withTags(Arrays.asList("foo", "bar")).withSecret(true),
            null, com.azure.core.util.Context.NONE);
    }
}
