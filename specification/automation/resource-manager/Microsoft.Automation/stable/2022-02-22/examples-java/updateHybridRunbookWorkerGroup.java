
import com.azure.resourcemanager.automation.models.HybridRunbookWorkerGroup;
import com.azure.resourcemanager.automation.models.RunAsCredentialAssociationProperty;

/**
 * Samples for HybridRunbookWorkerGroup Update.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2022-02-22/examples/
     * updateHybridRunbookWorkerGroup.json
     */
    /**
     * Sample code: Update hybrid worker group.
     * 
     * @param manager Entry point to AutomationManager.
     */
    public static void updateHybridWorkerGroup(com.azure.resourcemanager.automation.AutomationManager manager) {
        HybridRunbookWorkerGroup resource = manager.hybridRunbookWorkerGroups()
            .getWithResponse("rg", "testaccount", "TestHybridGroup", com.azure.core.util.Context.NONE).getValue();
        resource.update()
            .withCredential(new RunAsCredentialAssociationProperty().withName("myRunAsCredentialUpdatedName")).apply();
    }
}
