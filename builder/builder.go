package builder

import "html/template"

// В реализации ориентируемся на https://developer.android.com/guide/topics/ui/declaring-layout.html
// Строим из блоков:
// Группа - универсальная, группирует элементы, может принадлежать группе страниц
// Группа страниц - заголовок для подчиненных групп-страниц
// Командная панель - панель с кнопками, подменю, группами кнопок
// Поле - поле с текстом и т.п., редактируемое
// Кнопка - кнопка
// Таблица - таблица с данными
// TextView - Надпись - текстовое нередактируемое
// Картинка - картинка
// Панель навигации - перечень команд сбоку

type (
	Context interface {
		NextId() string // сдедующий уникальный id элемента в контексте, строкой
	}

	View interface {
		Render(Context) (template.HTML, error) // возвращает HTML
		Parent() ViewGroup
	}

	ViewGroup interface {
		View
		Items() []View
		AddItem(View)
	}
)
