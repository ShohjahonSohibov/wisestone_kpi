
-- Creating the kpi_tracking table
CREATE TABLE kpi_tracking (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    employee_id UUID NOT NULL,
    kpi_sub_sub_category_id UUID NOT NULL,
    month DATE NOT NULL,
    progress_percentage DECIMAL(5,2),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    FOREIGN KEY (employee_id) REFERENCES employees(id) ON DELETE CASCADE,
    FOREIGN KEY (kpi_sub_sub_category_id) REFERENCES kpi_sub_sub_categories(id) ON DELETE CASCADE
);
