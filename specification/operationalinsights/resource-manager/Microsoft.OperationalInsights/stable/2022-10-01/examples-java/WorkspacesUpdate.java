
import com.azure.resourcemanager.loganalytics.models.Workspace;
import com.azure.resourcemanager.loganalytics.models.WorkspaceCapping;
import com.azure.resourcemanager.loganalytics.models.WorkspaceSku;
import com.azure.resourcemanager.loganalytics.models.WorkspaceSkuNameEnum;

/**
 * Samples for Workspaces Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2022-10-01/examples/
     * WorkspacesUpdate.json
     */
    /**
     * Sample code: WorkspacesPatch.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void workspacesPatch(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        Workspace resource = manager.workspaces()
            .getByResourceGroupWithResponse("oiautorest6685", "oiautorest6685", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withSku(new WorkspaceSku().withName(WorkspaceSkuNameEnum.PER_GB2018)).withRetentionInDays(30)
            .withWorkspaceCapping(new WorkspaceCapping().withDailyQuotaGb(-1.0D)).apply();
    }
}
